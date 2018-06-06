// Copyright 2017-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package awsutils

// InstanceENIsAvailable contains a mapping of instance types to the number of ENIs available which is described at
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI
var InstanceENIsAvailable = map[string]int{
        "c1.medium":    2,
        "c1.xlarge":    4,
        "c3.large":     3,
        "c3.xlarge":    4,
        "c3.2xlarge":   4,
        "c3.4xlarge":   8,
        "c3.8xlarge":   8,
        "c4.large":     3,
        "c4.xlarge":    4,
        "c4.2xlarge":   4,
        "c4.4xlarge":   8,
        "c4.8xlarge":   8,
        "c5.large":     3,
        "c5.xlarge":    4,
        "c5.2xlarge":   4,
        "c5.4xlarge":   8,
        "c5.9xlarge":   8,
        "c5.18xlarge":  15,
        "c5d.large":    3,
        "c5d.xlarge":   4,
        "c5d.2xlarge":  4,
        "c5d.4xlarge":  8,
        "c5d.9xlarge":  8,
        "c5d.18xlarge": 15,
        "cc2.8xlarge":  8,
        "cr1.8xlarge":  8,
        "d2.xlarge":    4,
        "d2.2xlarge":   4,
        "d2.4xlarge":   8,
        "d2.8xlarge":   8,
        "f1.2xlarge":   4,
        "f1.16xlarge":  8,
        "g2.2xlarge":   4,
        "g2.8xlarge":   8,
        "g3.4xlarge":   8,
        "g3.8xlarge":   8,
        "g3.16xlarge":  15,
        "h1.2xlarge":   4,
        "h1.4xlarge":   8,
        "h1.8xlarge":   8,
        "h1.16xlarge":  15,
        "hs1.8xlarge":  8,
        "i2.xlarge":    4,
        "i2.2xlarge":   4,
        "i2.4xlarge":   8,
        "i2.8xlarge":   8,
        "i3.large":     3,
        "i3.xlarge":    4,
        "i3.2xlarge":   4,
        "i3.4xlarge":   8,
        "i3.8xlarge":   8,
        "i3.16xlarge":  15,
        "i3.metal":     15,
        "m1.small":     2,
        "m1.medium":    2,
        "m1.large":     3,
        "m1.xlarge":    4,
        "m2.xlarge":    4,
        "m2.2xlarge":   4,
        "m2.4xlarge":   8,
        "m3.medium":    2,
        "m3.large":     3,
        "m3.xlarge":    4,
        "m3.2xlarge":   4,
        "m4.large":     2,
        "m4.xlarge":    4,
        "m4.2xlarge":   4,
        "m4.4xlarge":   8,
        "m4.10xlarge":  8,
        "m4.16xlarge":  8,
        "m5.large":     3,
        "m5.xlarge":    4,
        "m5.2xlarge":   4,
        "m5.4xlarge":   8,
        "m5.12xlarge":  8,
        "m5.24xlarge":  15,
        "m5d.large":    3,
        "m5d.xlarge":   4,
        "m5d.2xlarge":  4,
        "m5d.4xlarge":  8,
        "m5d.12xlarge": 8,
        "m5d.24xlarge": 15,
        "p2.xlarge":    4,
        "p2.8xlarge":   8,
        "p2.16xlarge":  8,
        "p3.2xlarge":   4,
        "p3.8xlarge":   8,
        "p3.16xlarge":  8,
        "r3.large":     3,
        "r3.xlarge":    4,
        "r3.2xlarge":   4,
        "r3.4xlarge":   8,
        "r3.8xlarge":   8,
        "r4.large":     3,
        "r4.xlarge":    4,
        "r4.2xlarge":   4,
        "r4.4xlarge":   8,
        "r4.8xlarge":   8,
        "r4.16xlarge":  15,
        "t1.micro":     2,
        "t2.nano":      2,
        "t2.micro":     2,
        "t2.small":     2,
        "t2.medium":    3,
        "t2.large":     3,
        "t2.xlarge":    3,
        "t2.2xlarge":   3,
        "x1.16xlarge":  8,
        "x1.32xlarge":  8,
        "x1e.xlarge":   3,
        "x1e.2xlarge":  4,
        "x1e.4xlarge":  4,
        "x1e.8xlarge":  4,
        "x1e.16xlarge": 8,
        "x1e.32xlarge": 8,
}

// InstanceIPsAvailable contains a mapping of instance types to the number of IPs per ENI
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI
var InstanceIPsAvailable = map[string]int64{
        "c1.medium":    6,
        "c1.xlarge":    15,
        "c3.large":     10,
        "c3.xlarge":    15,
        "c3.2xlarge":   15,
        "c3.4xlarge":   30,
        "c3.8xlarge":   30,
        "c4.large":     10,
        "c4.xlarge":    15,
        "c4.2xlarge":   15,
        "c4.4xlarge":   30,
        "c4.8xlarge":   30,
        "c5.large":     10,
        "c5.xlarge":    15,
        "c5.2xlarge":   15,
        "c5.4xlarge":   30,
        "c5.9xlarge":   30,
        "c5.18xlarge":  50,
        "c5d.large":    10,
        "c5d.xlarge":   15,
        "c5d.2xlarge":  15,
        "c5d.4xlarge":  30,
        "c5d.9xlarge":  30,
        "c5d.18xlarge": 50,
        "cc2.8xlarge":  30,
        "cr1.8xlarge":  30,
        "d2.xlarge":    15,
        "d2.2xlarge":   15,
        "d2.4xlarge":   30,
        "d2.8xlarge":   30,
        "f1.2xlarge":   15,
        "f1.16xlarge":  50,
        "g2.2xlarge":   15,
        "g2.8xlarge":   30,
        "g3.4xlarge":   30,
        "g3.8xlarge":   30,
        "g3.16xlarge":  50,
        "h1.2xlarge":   15,
        "h1.4xlarge":   30,
        "h1.8xlarge":   30,
        "h1.16xlarge":  50,
        "hs1.8xlarge":  30,
        "i2.xlarge":    15,
        "i2.2xlarge":   15,
        "i2.4xlarge":   30,
        "i2.8xlarge":   30,
        "i3.large":     10,
        "i3.xlarge":    15,
        "i3.2xlarge":   15,
        "i3.4xlarge":   30,
        "i3.8xlarge":   30,
        "i3.16xlarge":  50,
        "i3.metal":     50,
        "m1.small":     4,
        "m1.medium":    6,
        "m1.large":     10,
        "m1.xlarge":    15,
        "m2.xlarge":    15,
        "m2.2xlarge":   30,
        "m2.4xlarge":   30,
        "m3.medium":    6,
        "m3.large":     10,
        "m3.xlarge":    15,
        "m3.2xlarge":   30,
        "m4.large":     10,
        "m4.xlarge":    15,
        "m4.2xlarge":   15,
        "m4.4xlarge":   30,
        "m4.10xlarge":  30,
        "m4.16xlarge":  30,
        "m5.large":     10,
        "m5.xlarge":    15,
        "m5.2xlarge":   15,
        "m5.4xlarge":   30,
        "m5.12xlarge":  30,
        "m5.24xlarge":  50,
        "m5d.large":    10,
        "m5d.xlarge":   15,
        "m5d.2xlarge":  15,
        "m5d.4xlarge":  30,
        "m5d.12xlarge": 30,
        "m5d.24xlarge": 50,
        "p2.xlarge":    15,
        "p2.8xlarge":   30,
        "p2.16xlarge":  30,
        "p3.2xlarge":   15,
        "p3.8xlarge":   30,
        "p3.16xlarge":  30,
        "r3.large":     10,
        "r3.xlarge":    15,
        "r3.2xlarge":   15,
        "r3.4xlarge":   30,
        "r3.8xlarge":   30,
        "r4.large":     10,
        "r4.xlarge":    15,
        "r4.2xlarge":   15,
        "r4.4xlarge":   30,
        "r4.8xlarge":   30,
        "r4.16xlarge":  50,
        "t1.micro":     2,
        "t2.nano":      2,
        "t2.micro":     2,
        "t2.small":     4,
        "t2.medium":    6,
        "t2.large":     12,
        "t2.xlarge":    15,
        "t2.2xlarge":   15,
        "x1.16xlarge":  30,
        "x1.32xlarge":  30,
        "x1e.xlarge":   10,
        "x1e.2xlarge":  15,
        "x1e.4xlarge":  15,
        "x1e.8xlarge":  15,
        "x1e.16xlarge": 30,
        "x1e.32xlarge": 30,
}
