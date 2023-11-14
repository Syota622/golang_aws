# # Lambda Function
# resource "aws_lambda_function" "this" {
#   filename      = "your-lambda-zip-package.zip"
#   function_name = "your-lambda-function-name"
#   role          = "your-lambda-execution-role-arn"
#   handler       = "index.handler" # depending on your code
#   runtime       = "nodejs14.x"
# }