<html>
    <head>
    <title></title>
    </head>
    <body>
        <h1>User Registration</h1>
        <h2 style="color:red">{{.Error}}</h2>
        <form action="/register" method="post">
            Email:<input type="email" name="username">
            First Name:<input type="text" name="firstname">
            Last Name:<input type="text" name="lastname">
            Password:<input type="password" name="password">
            <input type="submit" value="Register">
        </form>
    </body>
</html>