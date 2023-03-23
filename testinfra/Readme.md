## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | ~> 3.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_eip.fmc-mgmt-eip](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip) | resource |
| [aws_eip_association.fmc-mgmt-ip-assocation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip_association) | resource |
| [aws_instance.fmcv](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance) | resource |
| [aws_internet_gateway.int_gw](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway) | resource |
| [aws_network_interface.mgmt](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface) | resource |
| [aws_network_interface_sg_attachment.fmc_mgmt_attachment](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface_sg_attachment) | resource |
| [aws_route.ext_default_route](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route) | resource |
| [aws_route_table.fmc_mgmt_route](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table) | resource |
| [aws_route_table_association.mgmt_association](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association) | resource |
| [aws_security_group.allow_all](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_subnet.mgmt](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet) | resource |
| [aws_vpc.fmc_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |
| [aws_ami.fmcv](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami) | data source |
| [aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones) | data source |
| [aws_vpc.selected](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/vpc) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_azs"></a> [azs](#input\_azs) | AWS Availability Zones | `list` | `[]` | no |
| <a name="input_fmc_size"></a> [fmc\_size](#input\_fmc\_size) | Size of the FMCv instance | `string` | `"c5.4xlarge"` | no |
| <a name="input_fmc_version"></a> [fmc\_version](#input\_fmc\_version) | Version of the FMCv | `string` | `"fmcv-7.0.0"` | no |
| <a name="input_hostname"></a> [hostname](#input\_hostname) | FMCv OS hostname | `string` | `"fmc"` | no |
| <a name="input_igw_id"></a> [igw\_id](#input\_igw\_id) | Existing Internet Gateway ID | `string` | `""` | no |
| <a name="input_instances"></a> [instances](#input\_instances) | Number of FMCv instances | `number` | `1` | no |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | AWS EC2 Key | `any` | n/a | yes |
| <a name="input_name_tag_prefix"></a> [name\_tag\_prefix](#input\_name\_tag\_prefix) | Prefix for the 'Name' tag of the resources | `string` | `"FMCv"` | no |
| <a name="input_password"></a> [password](#input\_password) | Password for FMCv | `string` | `"P@$$w0rd1234"` | no |
| <a name="input_subnet_size"></a> [subnet\_size](#input\_subnet\_size) | Size of Management subnet | `number` | `24` | no |
| <a name="input_subnets"></a> [subnets](#input\_subnets) | mgmt subnets | `list` | `[]` | no |
| <a name="input_vpc_cidr"></a> [vpc\_cidr](#input\_vpc\_cidr) | VPC CIDR | `string` | `"10.0.0.0/16"` | no |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | Existing VPC ID | `string` | `""` | no |
| <a name="input_vpc_name"></a> [vpc\_name](#input\_vpc\_name) | VPC Name | `string` | `"Cisco-FMCv"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_FMCv_EIPs"></a> [FMCv\_EIPs](#output\_FMCv\_EIPs) | n/a |
