package main

const indexTemplate string = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>DDNS</title>
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.1.0/css/font-awesome.min.css">
    <!-- Optional theme -->
    <style type="text/css" media="all">
    
    /* Space out content a bit */
    body {
        padding-top: 20px;
        padding-bottom: 20px;
        background-color: #020402;
    }
    
    /* Everything but the jumbotron gets side spacing for mobile first views */
    .header,
    .marketing,
    .footer {
        padding-right: 15px;
        padding-left: 15px;
    }
    
    /* Custom page header */
    .header {
        border-bottom: 0px solid #54544c;
    }
    
    /* Make the masthead heading the same height as the navigation */
    .header h3 {
        padding-bottom: 8px;
        margin-top: 0;
        margin-bottom: 0;
        line-height: 40px;
        color: #b5b19d;
    }
    
    /* Set header buttons with theme */
    .nav-pills>li.active>a {
        background-color: #b5b19d;
    }
    .nav-pills>li.active>a:hover,
    .nav-pills>li.active>a:focus {
        color: #fff;
        background-color: #ede9d8;
    }
    .nav>li>a:hover,
    .nav>li>a:focus {
        background-color: #000;
        color: #ede9d8;
    }

    /* Custom page footer */
    .footer {
        padding-top: 19px;
        color: #777;
        border-top: 1px solid #54544c;
    }
    
    /* Customize container */
    @media (min-width: 768px) {
        .container {
            max-width: 730px;
        }
    }

    .container-narrow>hr {
        margin: 30px 0;
    }
    
    /* Description and button */
    .jumbotron {
        text-align: center;
        border-bottom: 1px solid #61676b;
        background-color: #2b2b27;
        color: #ede9d8;
    }
    .jumbotron .btn {
        padding: 14px 24px;
        margin-top: 5px;
        font-size: 21px;
        color: #020402;
        background-color: #ede9d8;
        border-color: #ede9d8;
    }
    
    /* Input box for hostname */
    .form-control {
        color: #555;
        background-color: #eee;
        border: 2px solid #61676b;
        border-radius: 5px;
        border-color: #eee;
    }
    .input-group-addon {
        color: #555;
        text-align: center;
        background-color: #fff;
        border: 1px solid #ccc;
        border-radius: 4px;
    }
    
    /* Responsive: Portrait tablets and up */
    @media screen and (min-width: 768px) {
        /* Remove the padding we set earlier */
        .header,
        .marketing,
        .footer {
            padding-right: 0;
            padding-left: 0;
        }
        /* Space out the masthead */
        .header {
            margin-bottom: 30px;
        }
        /* Remove the bottom border on the jumbotron for visual effect */
        .jumbotron {
            border-bottom: 0;
        }
    }
    </style>
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
        <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
        <![endif]-->
</head>

<body>
    <div class="container">
        <div class="header">
        </div>
        <div class="jumbotron">
            <h1>LeetDNS</h1>
            <p class="lead">The Dynamic DNS service that doesn't censor.
                <br>Like DuckDNS, but less autistic.</p>
            <hr style='background-color:#54544c;border-width:0;color:#54544c;height:1px;line-height:0;' />
            <form class="form-inline" role="form">
                <div id="hostname_input" class="form-group">
                    <div class="input-group">
                        <input id="hostname" class="form-control input-lg" type="text" placeholder="hostname">
                        <div class="input-group-addon input-lg">{{.domain}}</div>
                    </div>
                </div>
            </form>
            <hr style='background-color:#54544c;border-width:0;color:#54544c;height:1px;line-height:0;' />
            <input type="button" id="register" class="btn btn-primary disabled" value="Register Host" />
        </div>
        <div id="command_output"></div>
        <div class="footer">
            <p>Credit to pboehm for the original <a href="https://github.com/pboehm/ddns" style="color: #61676b;">source code.</a></p>
        </div>
    </div>
    <!-- /container -->
    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
    <!-- Latest compiled and minified JavaScript -->
    <script src="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
    <script type="text/javascript" charset="utf-8">
    function isValid() {
        $('#register').removeClass("disabled");
        $('#hostname_input').removeClass("has-error");
        $('#hostname_input').addClass("has-success");
    }

    function isNotValid(argument) {
        $('#register').addClass("disabled");
        $('#hostname_input').removeClass("has-success");
        $('#hostname_input').addClass("has-error");
    }

    function validate() {
        var hostname = $('#hostname').val();

        $.getJSON("/available/" + hostname, function(data) {
            if (data.available) {
                isValid();
            } else {
                isNotValid();
            }
        }).error(function() { isNotValid(); });
    }

    $(document).ready(function() {
        var timer = null;
        $('#hostname').on('keydown', function() {
            clearTimeout(timer);
            timer = setTimeout(validate, 800)
        });


        $('#register').click(function() {
            var hostname = $("#hostname").val();

            $.getJSON("/new/" + hostname, function(data) {
                console.log(data);

                var host = location.protocol + '//' + location.host;

                $("#command_output").append(
                    "<pre>curl \"" + host +
                    data.update_link + "\"</pre>");
            })
        });
    });
    </script>
</body>

</html>
`
