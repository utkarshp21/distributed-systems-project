<html>
    <head>
    <title></title>
    </head>
    <body>
        <h1>Hi User</h1>
        <form action="/follow" method="post">
            Email id:<input type="text" name="username">
            <input type="submit" value="Follow">
        </form>
         <form action="/tweet" method="post">
            Enter text:<input type="text" name="tweet">
            <input type="submit" value="Tweet">
          </form>
         <form action="/feed" method="post">
            <input type="submit" value="Feed">
          </form>
         <form action="/signout" method="post">
            <input type="submit" value="Signout">
          </form>
    </body>
</html>