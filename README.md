
<p align="center">
<img src="https://github.com/elaugier/ApiGo/blob/master/apigo-logo-180.png?raw=true"/>
</p>

* * *

# IMPORTANT

ApiGo is currently in development. For some private reasons, i can't work hard on this project. I hope have more time in several months. Sorry for that.

* * *

## Table of Contents

1. [What is ApiGo?](#what-is-coreapi)
1. [Implementation](#implementation)
1. [Installation](#installation)
1. [Usage](#usage)
1. [Security](#security)
1. [API](#api)
1. [Contributing](#contributing)
1. [License](#license)


# What is ApiGo?

ApiGo is an API engine which allow to build quickly a REST API without development. If you have already some scripts written in Powershell, Python, ... and you need to expose these through an web API, ApiGo may be your solution.

![Overview](https://github.com/elaugier/ApiGo/blob/master/docs/apigo-overview.png)


## Script Languages Compliance

- [Powershell](https://docs.microsoft.com/powershell/)
- [Python](https://www.python.org/)
- [Perl](https://www.perl.org/)
- [Php](http://php.net/)
- [Ruby](https://www.ruby-lang.org/)

# Implementation

## Components

There are some components in ApiGo:

*; an API Engine: apigo-engine
*; a Worker: apigo-worker
*; a Message Broker: Apache Kafka, RabbitMQ, ZeroMQ*
*; a Database: Postgresql, MySQL, etc.

And as soon as possible:

*; an Authentication service
*; an Authorization service (per App Domain)

## Dependencies

* Messaging System

	- Apache Kafka : [Download](https://kafka.apache.org/downloads) [Quickstart](https://kafka.apache.org/quickstart)
  - Pivotal RabbitMQ : [Get Started](http://www.rabbitmq.com/#getstarted)
  - ZeroMQ : [Learn Basics](http://zeromq.org/intro:read-the-manual) [Download](http://zeromq.org/intro:get-the-software)

* Database compliance

	- Postgresql
	- MySQL
	- Sql Server
	- Sqlite

# Installation

Before install the components, you must build the project.

## Build and publish ApiGo

Open a command line in __ApiGo__ git root folder and type the following command : 

~~~~

Linux : TBD

Windows : build.cmd

~~~~

**[Back to top](#table-of-contents)**

## ApiGo installation

TBD

**[Back to top](#table-of-contents)**

## Setup ApiGo as windows service

To setup coreApi as windows service, we recommand to use [NSSM](https://nssm.cc/download)

Follow the instructions written on the following url [https://nssm.cc/usage](https://nssm.cc/usage)

**[Back to top](#table-of-contents)**

# Usage

## Engine Configuration

All configurations are stored in the folder __config__.

At root, there is a file _default.json_. This is the main configuration.

~~~~
{
  "AccountingDatabase": {
    "AdminDatabase": "",
    "ConnectionString": "",
    "Driver": "postgres"
  },
  "Bindings": "0.0.0.0:1203",
  "CertPath": "",
  "CertPwd": "",
  "JobsDatabase": {
    "AdminDatabase": "",
    "ConnectionString": "",
    "Driver": "postgres"
  },
  "KafkaProducer": {
    "BootstrapServers": "localhost:9092"
  },
  "MaxConcurrentConnections": "0",
  "MaxConcurrentUpgradeConnections": "0",
  "MaxRequestBodySize": "338657280",
  "RoutesConfigPath": "config/routes",
  "Secure": "false"
}
~~~~

For each route you have to create two files. First with __.conf.json__ like as following for a Powershell Cmdlet: 

~~~~
{
  "Name": "Route1",
  "Cmd": {
    "Name": "Command1",
    "Type": "Powershell",
    "PSModule": "PSModule",
    "Params": [
      {
        "Name": "Argument1",
        "Type": "String",
        "Mandatory": "True",
        "In": "body"
      }
    ]
  },
  "Route": "/route1",
  "Method": "POST",
  "JobType": "synchronous",
  "Topic": "topic1",
  "Timeout": "15000",
  "AddRequestIdParam": "True"
}
~~~~

  * __Type__ : must be equal to "Powershell" or "Python" or "Perl" or "CommandLine"
  * __PSModule__ : only for Powershell type, you must specify the module where the cmdlet is defined
  * __Method__ : must be equal to "GET" or "POST" or "PUT" or "PATCH" or "DELETE"
  * __JobType__ : must be equal to "synchronous" or "asynchronous"
  * __Timeout__ : for synchronous job, the timeout is the duration until CoreApi considers that the job won't be completed
  * __Topic__ : this is the topic where all jobs for this API entry were sent

Same example for Python script:

~~~~
{
  "Name": "Route2",
  "Cmd": {
    "Name": "pyCommand1",
    "Type": "Python",
    "PyVenv": "venv1",
    "Params": [
      {
        "Name": "Argument1",
        "Type": "String",
        "Mandatory": "True",
        "In": "body"
      }
    ]
  },
  "Route": "/route2",
  "Method": "POST",
  "JobType": "synchronous",
  "Topic": "topic1",
  "Timeout": "15000",
  "AddRequestIdParam": "True"
}
~~~~

For Python script, you must define a new attribute "PyEnv" to allow the worker activate the good virtual environment

**[Back to top](#table-of-contents)**

## Worker Configuration

All configurations are stored in the folder __config__.

At root, there is a file _default.json_. This is the main configuration.

~~~~
{
  "AccountingDatabase": {
    "AdminDatabase": "",
    "ConnectionString": "",
    "Driver": "postgres"
  },
  "JobsDatabase": {
    "AdminDatabase": "",
    "ConnectionString": "",
    "Driver": "postgres"
  },
  "KafkaConsumer": {
    "BootstrapServers": "",
    "Debug": "",
    "GroupId": "winworkers"
  },
  "ScriptsPath": "config\\scripts",
  "WorkerTopic": "winworkersTopic",
  "OnMessageTimeout": "10000",
  "MaxConcurrentJobs":"0",
  "PoolWaitingTime","100"
}
~~~~

**[Back to top](#table-of-contents)**

## Worker scripts installation

### Overview

the ApiGo worker can use any command, you can launch from the OS shell. But you have to know how worker returns 
the result to the engine.

The engine return result as 'application/json' and the standard response have the following structure:

~~~~~
{
	"sts":"",
	"hco":"",
	"msg":"",
	"dbg":"",
	"dta":""
}
~~~~~

 * __sts__ : can contains these following values : "success", "failed"
 * __hco__ : contains the http code status
 * __msg__ : contains any useful message if needed
 * __dbg__	: contains debug informations if needed
 * __dta__ : contains string or any valid JSON structure returned by the command executed by the worker 

**[Back to top](#table-of-contents)**

### Powershell CmdLets and Modules

To use Powershell with ApiGo, you must use Powershell Module (.psm1), not directly a script (.ps1).  The main interest of the module is to limit the number of scripts to deploy on the workers by grouping the cmdlets in a single module file.

The following code is an minimal example of Powershell module.

```powershell

# ApiGo (https://github.com/elaugier/ApiGo)
# -----------------------------------------
# script sample for Powershell

function HelloWorld {
    param (
    )
    @{ "message" = "Hello World!" } | ConvertTo-Json
}

Export-ModuleMember -Function HelloWorld

```

To make a PS Module available for ApiGo, you just have to copy your file under the folder __psmodules__ in the scripts directory (watch the worker configuration file to know it)

**[Back to top](#table-of-contents)**

### Python scripts and virtual environments

To use Python scripts, you need to configure at least one virtual environment. 

1. Open a shell under the folder __PyEnvs__ in the scripts directory (watch the worker configuration file to know it)
1. Create a virtual environment with the following command :
~~~~
        python -m venv [VirtualEnvName]
~~~~
1. To make a Python script available for ApiGo, you just have to copy your file under the folder __python-scripts__ in the scripts directory

The following code is an minimal example of Python script.

```python

# ApiGo (https://github.com/elaugier/ApiGo)
# -----------------------------------------
# script sample for Python

import json

t = {'message':'Hello World'}
print(json.dumps(t))

```

**[Back to top](#table-of-contents)**

### Perl

To make a Perl script available for ApiGo, you just have to copy your file under the folder __perl-scripts__ in the scripts directory (watch the worker configuration file to know it)

The following code is an minimal example of Perl script.

```perl

# ApiGo (https://github.com/elaugier/ApiGo)
# -----------------------------------------
# script sample for Perl

use JSON;

my %msg = ('message' => 'Hello World!');
my $json = encode_json \%msg;
print $json

```

**[Back to top](#table-of-contents)**

### Ruby

To make a Ruby script available for ApiGo, you just have to copy your file under the folder __ruby-scripts__ in the scripts directory (watch the worker configuration file to know it)

The following code is an minimal example of Ruby script.

```ruby

# ApiGo (https://github.com/elaugier/ApiGo)
# -----------------------------------------
# script sample for Ruby

require 'json/ext'

t = {:message => "Hello World!"}
puts t.to_json

```

**[Back to top](#table-of-contents)**

### PHP

To make a PHP script available for ApiGo, you just have to copy your file under the folder __php-scripts__ in the scripts directory (watch the worker configuration file to know it)

The following code is an minimal example of PHP script.

```php

<?php
/**
 * ApiGo (https://github.com/elaugier/ApiGo)
 * -----------------------------------------
 * 
 * script sample for PHP
 */
$t = [
    "message" => "Hello World!"
];

echo json_encode($t);

```


**[Back to top](#table-of-contents)**

# Security

ApiGo only supports the Key API for the moment. It will evolve in the future. 

# API

**[Back to top](#table-of-contents)**

# Contributing

Open an issue first to discuss potential changes/additions.

**[Back to top](#table-of-contents)**

# License

GNU General Public License v3.0

[![Egen Guru logo](https://github.com/elaugier/ApiGo/blob/master/docs/g887.png)](https://egen.guru/)

contact : support@egen.guru

**[Back to top](#table-of-contents)**
