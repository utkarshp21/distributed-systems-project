<html>
    <head>
    <title></title>
    </head>
    <body>
        <h1>Hi User</h1>
        <h2 style="color:red">{{.Error}}</h2>
        <h2 style="color: green">{{.Success}}</h2>
        <form action="/follow" method="post">
            Email id:<input type="text" name="username">
            <input type="submit" value="Follow">
        </form>
        <form action="/unfollow" method="post">
            Email id:<input type="text" name="username">
            <input type="submit" value="Unfollow">
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
          <p>
            {{.Feed}}
          </p>
    </body>
</html>