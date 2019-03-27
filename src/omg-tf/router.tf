/***********
 * ERT LB  *
 ***********/

resource "google_compute_firewall" "cf-public" {
  name    = "${var.env_name}-cf-public"
  network = "${google_compute_network.pcf-network.name}"

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags = ["${var.env_name}-httpslb", "${var.env_name}-isoseglb"]
}

resource "google_compute_global_address" "cf" {
  name = "${var.env_name}-cf"
}

resource "google_compute_instance_group" "httplb" {
  // Count based on number of AZs
  count       = 3
  name        = "${var.env_name}-httpslb-${element(var.zones, count.index)}"
  description = "Instance Group for ERT HTTP Routers"
  zone        = "${element(var.zones, count.index)}"
}

resource "google_compute_http_health_check" "cf-public" {
  name                = "${var.env_name}-cf-public"
  port                = 8080
  request_path        = "/health"
  check_interval_sec  = 5
  timeout_sec         = 3
  healthy_threshold   = 6
  unhealthy_threshold = 3
}

resource "google_compute_backend_service" "http_lb_backend_service" {
  name        = "${var.env_name}-httpslb"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 900
  enable_cdn  = false

  backend {
    group = "${google_compute_instance_group.httplb.0.self_link}"
  }

  backend {
    group = "${google_compute_instance_group.httplb.1.self_link}"
  }

  backend {
    group = "${google_compute_instance_group.httplb.2.self_link}"
  }

  health_checks = ["${google_compute_http_health_check.cf-public.self_link}"]
}

resource "google_compute_url_map" "https_lb_url_map" {
  name = "${var.env_name}-cf-http"

  default_service = "${google_compute_backend_service.http_lb_backend_service.self_link}"
}

resource "google_compute_target_http_proxy" "http_lb_proxy" {
  name        = "${var.env_name}-httpproxy"
  description = "Proxy public HTTP traffic"
  url_map     = "${google_compute_url_map.https_lb_url_map.self_link}"
}

resource "google_compute_target_https_proxy" "https_lb_proxy" {
  name             = "${var.env_name}-httpsproxy"
  description      = "Proxy public HTTPS traffic"
  url_map          = "${google_compute_url_map.https_lb_url_map.self_link}"
  ssl_certificates = ["${google_compute_ssl_certificate.cert.self_link}"]
}

resource "google_compute_ssl_certificate" "cert" {
  name_prefix = "${var.env_name}-lbcert"
  description = "user provided ssl private key / ssl certificate pair"
  private_key = "${var.ssl_cert_private_key}"
  certificate = "${var.ssl_cert}"

  lifecycle = {
    create_before_destroy = true
  }
}

resource "google_compute_firewall" "cf-health_check" {
  name    = "${var.env_name}-cf-health-check"
  network = "${google_compute_network.pcf-network.name}"

  allow {
    protocol = "tcp"
    ports    = ["8080"]
  }

  source_ranges = ["130.211.0.0/22", "35.191.0.0/16"]
  target_tags   = ["${var.env_name}-httpslb", "${var.env_name}-isoseglb"]
}

resource "google_compute_global_forwarding_rule" "cf-http" {
  name       = "${var.env_name}-cf-lb-http"
  ip_address = "${google_compute_global_address.cf.address}"
  target     = "${google_compute_target_http_proxy.http_lb_proxy.self_link}"
  port_range = "80"
}

resource "google_compute_global_forwarding_rule" "cf-https" {
  name       = "${var.env_name}-cf-lb-https"
  ip_address = "${google_compute_global_address.cf.address}"
  target     = "${google_compute_target_https_proxy.https_lb_proxy.self_link}"
  port_range = "443"
}

/*****************
 * Diego SSH LB  *
 *****************/

resource "google_compute_firewall" "cf-ssh" {
  name    = "${var.env_name}-cf-ssh"
  network = "${google_compute_network.pcf-network.name}"

  allow {
    protocol = "tcp"
    ports    = ["2222"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["${var.env_name}-cf-ssh"]
}

resource "google_compute_address" "cf-ssh" {
  name = "${var.env_name}-cf-ssh"
}

resource "google_compute_target_pool" "cf-ssh" {
  name             = "${var.env_name}-cf-ssh"
  session_affinity = "NONE"
}

resource "google_compute_forwarding_rule" "cf-ssh" {
  name        = "${var.env_name}-cf-ssh"
  target      = "${google_compute_target_pool.cf-ssh.self_link}"
  port_range  = "2222-2222"
  ip_protocol = "TCP"
  ip_address  = "${google_compute_address.cf-ssh.address}"
}


/*****************
 * Web Socket LB *
 *****************/

resource "google_compute_firewall" "cf-wss" {
  name    = "${var.env_name}-cf-wss"
  network = "${google_compute_network.pcf-network.name}"

  allow {
    protocol = "tcp"
    ports    = ["443"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["${var.env_name}-cf-wss"]
}

resource "google_compute_address" "cf-wss" {
  name = "${var.env_name}-cf-wss"
}

resource "google_compute_target_pool" "cf-wss" {
  name             = "${var.env_name}-cf-wss"
  session_affinity = "NONE"
}

resource "google_compute_forwarding_rule" "cf-wss" {
  name        = "${var.env_name}-cf-wss"
  target      = "${google_compute_target_pool.cf-wss.self_link}"
  port_range  = "443-443"
  ip_protocol = "TCP"
  ip_address  = "${google_compute_address.cf-wss.address}"
}
