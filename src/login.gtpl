<html>
    <head>
    <title></title>
    </head>
    <body>
        <h1>Please Login</h1>
        <h2 style="color:red">{{.Error}}</h2>
        <form action="/login" method="post">
            Email:<input type="email" name="username" required>
            Password:<input type="password" name="password" required>
            <input type="submit" value="Login">
        </form>
    <button onclick=goTLogin()>Register</button>
    </body>
    <script>
        function goTLogin() {
            window.location.href = '/register';
        }
    </script>
</html>