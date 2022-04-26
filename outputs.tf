# Output value definitions

output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_s3_bucket.lambda_bucket.id
}

output "base_api_gateway_url" {
  description = "Base URL for API Gateway stage."

  value = aws_apigatewayv2_stage.lambda.invoke_url
}

output "cloudfront_domain_name" {
  description = "Cloudfront domain name"

  value = aws_cloudfront_distribution.main.domain_name
}

