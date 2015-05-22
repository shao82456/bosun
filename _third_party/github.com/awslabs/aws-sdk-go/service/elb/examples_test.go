package elb_test

import (
	"bosun.org/_third_party/github.com/awslabs/aws-sdk-go/aws"
	"bytes"
	"fmt"
	"time"

	"bosun.org/_third_party/github.com/awslabs/aws-sdk-go/aws/awserr"
	"bosun.org/_third_party/github.com/awslabs/aws-sdk-go/aws/awsutil"
	"bosun.org/_third_party/github.com/awslabs/aws-sdk-go/service/elb"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleELB_AddTags() {
	svc := elb.New(nil)

	params := &elb.AddTagsInput{
		LoadBalancerNames: []*string{ // Required
			aws.String("AccessPointName"), // Required
			// More values...
		},
		Tags: []*elb.Tag{ // Required
			&elb.Tag{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_ApplySecurityGroupsToLoadBalancer() {
	svc := elb.New(nil)

	params := &elb.ApplySecurityGroupsToLoadBalancerInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		SecurityGroups: []*string{ // Required
			aws.String("SecurityGroupId"), // Required
			// More values...
		},
	}
	resp, err := svc.ApplySecurityGroupsToLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_AttachLoadBalancerToSubnets() {
	svc := elb.New(nil)

	params := &elb.AttachLoadBalancerToSubnetsInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		Subnets: []*string{ // Required
			aws.String("SubnetId"), // Required
			// More values...
		},
	}
	resp, err := svc.AttachLoadBalancerToSubnets(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_ConfigureHealthCheck() {
	svc := elb.New(nil)

	params := &elb.ConfigureHealthCheckInput{
		HealthCheck: &elb.HealthCheck{ // Required
			HealthyThreshold:   aws.Long(1),                     // Required
			Interval:           aws.Long(1),                     // Required
			Target:             aws.String("HealthCheckTarget"), // Required
			Timeout:            aws.Long(1),                     // Required
			UnhealthyThreshold: aws.Long(1),                     // Required
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.ConfigureHealthCheck(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_CreateAppCookieStickinessPolicy() {
	svc := elb.New(nil)

	params := &elb.CreateAppCookieStickinessPolicyInput{
		CookieName:       aws.String("CookieName"),      // Required
		LoadBalancerName: aws.String("AccessPointName"), // Required
		PolicyName:       aws.String("PolicyName"),      // Required
	}
	resp, err := svc.CreateAppCookieStickinessPolicy(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_CreateLBCookieStickinessPolicy() {
	svc := elb.New(nil)

	params := &elb.CreateLBCookieStickinessPolicyInput{
		LoadBalancerName:       aws.String("AccessPointName"), // Required
		PolicyName:             aws.String("PolicyName"),      // Required
		CookieExpirationPeriod: aws.Long(1),
	}
	resp, err := svc.CreateLBCookieStickinessPolicy(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_CreateLoadBalancer() {
	svc := elb.New(nil)

	params := &elb.CreateLoadBalancerInput{
		Listeners: []*elb.Listener{ // Required
			&elb.Listener{ // Required
				InstancePort:     aws.Long(1),            // Required
				LoadBalancerPort: aws.Long(1),            // Required
				Protocol:         aws.String("Protocol"), // Required
				InstanceProtocol: aws.String("Protocol"),
				SSLCertificateID: aws.String("SSLCertificateId"),
			},
			// More values...
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
		AvailabilityZones: []*string{
			aws.String("AvailabilityZone"), // Required
			// More values...
		},
		Scheme: aws.String("LoadBalancerScheme"),
		SecurityGroups: []*string{
			aws.String("SecurityGroupId"), // Required
			// More values...
		},
		Subnets: []*string{
			aws.String("SubnetId"), // Required
			// More values...
		},
		Tags: []*elb.Tag{
			&elb.Tag{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_CreateLoadBalancerListeners() {
	svc := elb.New(nil)

	params := &elb.CreateLoadBalancerListenersInput{
		Listeners: []*elb.Listener{ // Required
			&elb.Listener{ // Required
				InstancePort:     aws.Long(1),            // Required
				LoadBalancerPort: aws.Long(1),            // Required
				Protocol:         aws.String("Protocol"), // Required
				InstanceProtocol: aws.String("Protocol"),
				SSLCertificateID: aws.String("SSLCertificateId"),
			},
			// More values...
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.CreateLoadBalancerListeners(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_CreateLoadBalancerPolicy() {
	svc := elb.New(nil)

	params := &elb.CreateLoadBalancerPolicyInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		PolicyName:       aws.String("PolicyName"),      // Required
		PolicyTypeName:   aws.String("PolicyTypeName"),  // Required
		PolicyAttributes: []*elb.PolicyAttribute{
			&elb.PolicyAttribute{ // Required
				AttributeName:  aws.String("AttributeName"),
				AttributeValue: aws.String("AttributeValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateLoadBalancerPolicy(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DeleteLoadBalancer() {
	svc := elb.New(nil)

	params := &elb.DeleteLoadBalancerInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.DeleteLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DeleteLoadBalancerListeners() {
	svc := elb.New(nil)

	params := &elb.DeleteLoadBalancerListenersInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		LoadBalancerPorts: []*int64{ // Required
			aws.Long(1), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteLoadBalancerListeners(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DeleteLoadBalancerPolicy() {
	svc := elb.New(nil)

	params := &elb.DeleteLoadBalancerPolicyInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		PolicyName:       aws.String("PolicyName"),      // Required
	}
	resp, err := svc.DeleteLoadBalancerPolicy(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DeregisterInstancesFromLoadBalancer() {
	svc := elb.New(nil)

	params := &elb.DeregisterInstancesFromLoadBalancerInput{
		Instances: []*elb.Instance{ // Required
			&elb.Instance{ // Required
				InstanceID: aws.String("InstanceId"),
			},
			// More values...
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.DeregisterInstancesFromLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DescribeInstanceHealth() {
	svc := elb.New(nil)

	params := &elb.DescribeInstanceHealthInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		Instances: []*elb.Instance{
			&elb.Instance{ // Required
				InstanceID: aws.String("InstanceId"),
			},
			// More values...
		},
	}
	resp, err := svc.DescribeInstanceHealth(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DescribeLoadBalancerAttributes() {
	svc := elb.New(nil)

	params := &elb.DescribeLoadBalancerAttributesInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.DescribeLoadBalancerAttributes(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DescribeLoadBalancerPolicies() {
	svc := elb.New(nil)

	params := &elb.DescribeLoadBalancerPoliciesInput{
		LoadBalancerName: aws.String("AccessPointName"),
		PolicyNames: []*string{
			aws.String("PolicyName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeLoadBalancerPolicies(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DescribeLoadBalancerPolicyTypes() {
	svc := elb.New(nil)

	params := &elb.DescribeLoadBalancerPolicyTypesInput{
		PolicyTypeNames: []*string{
			aws.String("PolicyTypeName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeLoadBalancerPolicyTypes(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DescribeLoadBalancers() {
	svc := elb.New(nil)

	params := &elb.DescribeLoadBalancersInput{
		LoadBalancerNames: []*string{
			aws.String("AccessPointName"), // Required
			// More values...
		},
		Marker:   aws.String("Marker"),
		PageSize: aws.Long(1),
	}
	resp, err := svc.DescribeLoadBalancers(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DescribeTags() {
	svc := elb.New(nil)

	params := &elb.DescribeTagsInput{
		LoadBalancerNames: []*string{ // Required
			aws.String("AccessPointName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTags(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DetachLoadBalancerFromSubnets() {
	svc := elb.New(nil)

	params := &elb.DetachLoadBalancerFromSubnetsInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		Subnets: []*string{ // Required
			aws.String("SubnetId"), // Required
			// More values...
		},
	}
	resp, err := svc.DetachLoadBalancerFromSubnets(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_DisableAvailabilityZonesForLoadBalancer() {
	svc := elb.New(nil)

	params := &elb.DisableAvailabilityZonesForLoadBalancerInput{
		AvailabilityZones: []*string{ // Required
			aws.String("AvailabilityZone"), // Required
			// More values...
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.DisableAvailabilityZonesForLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_EnableAvailabilityZonesForLoadBalancer() {
	svc := elb.New(nil)

	params := &elb.EnableAvailabilityZonesForLoadBalancerInput{
		AvailabilityZones: []*string{ // Required
			aws.String("AvailabilityZone"), // Required
			// More values...
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.EnableAvailabilityZonesForLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_ModifyLoadBalancerAttributes() {
	svc := elb.New(nil)

	params := &elb.ModifyLoadBalancerAttributesInput{
		LoadBalancerAttributes: &elb.LoadBalancerAttributes{ // Required
			AccessLog: &elb.AccessLog{
				Enabled:        aws.Boolean(true), // Required
				EmitInterval:   aws.Long(1),
				S3BucketName:   aws.String("S3BucketName"),
				S3BucketPrefix: aws.String("AccessLogPrefix"),
			},
			AdditionalAttributes: []*elb.AdditionalAttribute{
				&elb.AdditionalAttribute{ // Required
					Key:   aws.String("StringVal"),
					Value: aws.String("StringVal"),
				},
				// More values...
			},
			ConnectionDraining: &elb.ConnectionDraining{
				Enabled: aws.Boolean(true), // Required
				Timeout: aws.Long(1),
			},
			ConnectionSettings: &elb.ConnectionSettings{
				IdleTimeout: aws.Long(1), // Required
			},
			CrossZoneLoadBalancing: &elb.CrossZoneLoadBalancing{
				Enabled: aws.Boolean(true), // Required
			},
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.ModifyLoadBalancerAttributes(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_RegisterInstancesWithLoadBalancer() {
	svc := elb.New(nil)

	params := &elb.RegisterInstancesWithLoadBalancerInput{
		Instances: []*elb.Instance{ // Required
			&elb.Instance{ // Required
				InstanceID: aws.String("InstanceId"),
			},
			// More values...
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
	}
	resp, err := svc.RegisterInstancesWithLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_RemoveTags() {
	svc := elb.New(nil)

	params := &elb.RemoveTagsInput{
		LoadBalancerNames: []*string{ // Required
			aws.String("AccessPointName"), // Required
			// More values...
		},
		Tags: []*elb.TagKeyOnly{ // Required
			&elb.TagKeyOnly{ // Required
				Key: aws.String("TagKey"),
			},
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_SetLoadBalancerListenerSSLCertificate() {
	svc := elb.New(nil)

	params := &elb.SetLoadBalancerListenerSSLCertificateInput{
		LoadBalancerName: aws.String("AccessPointName"),  // Required
		LoadBalancerPort: aws.Long(1),                    // Required
		SSLCertificateID: aws.String("SSLCertificateId"), // Required
	}
	resp, err := svc.SetLoadBalancerListenerSSLCertificate(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_SetLoadBalancerPoliciesForBackendServer() {
	svc := elb.New(nil)

	params := &elb.SetLoadBalancerPoliciesForBackendServerInput{
		InstancePort:     aws.Long(1),                   // Required
		LoadBalancerName: aws.String("AccessPointName"), // Required
		PolicyNames: []*string{ // Required
			aws.String("PolicyName"), // Required
			// More values...
		},
	}
	resp, err := svc.SetLoadBalancerPoliciesForBackendServer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleELB_SetLoadBalancerPoliciesOfListener() {
	svc := elb.New(nil)

	params := &elb.SetLoadBalancerPoliciesOfListenerInput{
		LoadBalancerName: aws.String("AccessPointName"), // Required
		LoadBalancerPort: aws.Long(1),                   // Required
		PolicyNames: []*string{ // Required
			aws.String("PolicyName"), // Required
			// More values...
		},
	}
	resp, err := svc.SetLoadBalancerPoliciesOfListener(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}