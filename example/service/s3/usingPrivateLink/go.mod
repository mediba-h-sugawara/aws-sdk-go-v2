module github.com/aws/aws-sdk-go-v2/example/service/s3/usingPrivateLink

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.3.1
	github.com/aws/aws-sdk-go-v2/config v1.1.4
	github.com/aws/aws-sdk-go-v2/service/s3 v1.4.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.4.0
	github.com/aws/aws-sdk-go-v2/credentials v1.1.4
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.0.5
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.0.3
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.0.5
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sso v1.1.4
	github.com/aws/aws-sdk-go-v2/service/sts v1.2.1
)

replace github.com/aws/aws-sdk-go-v2/config => ../../../../config/

replace github.com/aws/aws-sdk-go-v2/service/s3 => ../../../../service/s3/

replace github.com/aws/aws-sdk-go-v2/service/s3control => ../../../../service/s3control/

replace github.com/aws/aws-sdk-go-v2 => ../../../../

replace github.com/aws/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/aws/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../../../../service/sts/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../../../service/internal/s3shared/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/sso => ../../../../service/sso/
