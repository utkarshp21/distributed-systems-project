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
          <p id="demo">
          </p>
          <style>
            table, th, td {
            padding: 10px;
            border: 1px solid black; 
            border-collapse: collapse;
            }
          </style>
          <script>
            var i;
            var table = "<table><tr><th>Username</th><th>Top 5 tweets</th></tr>";
            var list;
            var feed = {{.Feed}}.split('$')
            feed = feed.slice(0, feed.length - 1);
            for(i=0;i<feed.length;i++){
              list = feed[i].split(':')
              table += "<tr><td>"+list[0]+"</td><td>"+list[1].slice(0, list[1].length - 1)+"</td></tr>"
            }
            table += "</table>"
            document.getElementById("demo").innerHTML = table;
          </script>
    </body>
</html>