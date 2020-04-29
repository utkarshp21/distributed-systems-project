<html>
    <head>
    <title></title>
    </head>
    <body>
        <h1>Hi User</h1>
        <h2 style="color:red">{{.Error}}</h2>
        <h2 style="color: green">{{.Success}}</h2>
        <form action="/userlist" method="post">
            <input type="submit" value="List users">
        </form>
         <form action="/tweet" method="post">
            Enter text:<input type="text" name="tweet">
            <input type="submit" value="Tweet">
          </form>
         <form action="/feed" method="post">
            <input type="submit" value="Refresh feed">
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
            var j;
            var table = "";
            var list;
            var feed = {{.Feed}}.split('$');
            for(i=0;i<feed.length;i++){
              list = feed[i].split('^');
              table += "Top 5 tweets from : "+list[0]+"<br><br>";
              tweetlist = list[1].split(',');
              table += "<table><tr><th>Tweet</th><th>Timestamp</th></tr>"
              for(j=0;j<tweetlist.length;j++){
              	tweet = tweetlist[j].split('*');
              	table += "<tr><td>"+tweet[0]+"</td><td>"+tweet[1]+"</td></tr>"
                }
                table += "</table><br>"
              }
            document.getElementById("demo").innerHTML = table;
          </script>
    </body>
</html>