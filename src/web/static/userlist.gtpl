<!DOCTYPE html>
<html>
    <head>
    <title></title>
    </head>
    <body>
        <h2 style="color:red">{{.Error}}</h2>
        <h2 style="color: green">{{.Success}}</h2>
        <form action="/userlist" method="post">
            <input type="submit" value="Refresh list">
        </form>
        <p id="demo">
        </p>
        <form action="/follow" method="post">
            Email id:<input type="text" name="username">
            <input type="submit" value="Follow">
        </form>
        <br>
        <form action="/unfollow" method="post">
            Email id:<input type="text" name="username">
            <input type="submit" value="Unfollow">
         </form>
         <br>
         <button onclick=goTProfile()>Back</button>
        <style>
            table, th, td {
            padding: 10px;
            border: 1px solid black;
            border-collapse: collapse;
            }
        </style>
        <script>
            var i;
            var table = "<table><tr><th>User list</th></tr>";
            var lists = {{.List}}.split('$')
            userlist = lists[0].split(',')
            followlist = lists[1].split(',')
            for(i=0;i<userlist.length;i++){
              table += "<tr><td>"+userlist[i]+"</tr>"
            }
            table += "</table><br><table><tr><th>Followed users</th></tr>"
             for(i=0;i<followlist.length;i++){
              table += "<tr><td>"+followlist[i]+"</tr>"
            }
            table += "</table>"
            document.getElementById("demo").innerHTML = table;
        </script>
        <script>
                function goTProfile() {
                    window.location.href = '/feed';
                }
            </script>
    </body>
</html>


