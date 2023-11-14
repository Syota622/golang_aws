resource "aws_acm_certificate" "this" {
  domain_name       = "your-domain-name.com"
  validation_method = "DNS"

  tags = {
    Name = "${var.pj}-acm"
  }
}

# 証明書のバリデーション
resource "aws_acm_certificate_validation" "this" {
  certificate_arn         = aws_acm_certificate.this.arn
  validation_record_fqdns = [aws_route53_record.cert_validation.fqdn]
}