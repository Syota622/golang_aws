# パブリックサブネット
resource "aws_subnet" "public_subnet" {
  count = 3

  vpc_id                  = aws_vpc.main_vpc.id
  cidr_block              = "10.0.${count.index}.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1${element(["a", "c", "d"], count.index)}"

  tags = {
    Name = "${var.pj}-public-subnet-${element(["a", "c", "d"], count.index)}-${var.env}"
  }
}

# プライベートサブネット
resource "aws_subnet" "private_subnet" {
  count = 3

  vpc_id            = aws_vpc.main_vpc.id
  cidr_block        = "10.0.${count.index + 3}.0/24"
  availability_zone = "ap-northeast-1${element(["a", "c", "d"], count.index)}"

  tags = {
    Name = "${var.pj}-private-subnet-${element(["a", "c", "d"], count.index)}-${var.env}"
  }
}
