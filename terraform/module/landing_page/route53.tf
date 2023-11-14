# Route 53 ドメインのレコード設定
resource "aws_route53_record" "cert_validation" {
  name    = aws_acm_certificate.this.domain_validation_options.0.resource_record_name
  type    = aws_acm_certificate.this.domain_validation_options.0.resource_record_type
  zone_id = "your-zone-id"
  records = [aws_acm_certificate.this.domain_validation_options.0.resource_record_value]
  ttl     = 60
}