<html>
    <head>
    <title></title>
    </head>
    <body>
        <h1>Please Login</h1>
        <h2 style="color:red">{{.Error}}</h2>
        <form action="/login" method="post">
            Email:<input type="email" name="username">
            Password:<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>