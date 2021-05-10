document.addEventListener("DOMContentLoaded", function(){
    btn = document.getElementById("button")
    btn.onclick = function(){
        // get user search payload
        payload = document.getElementById("search").value;
        
        // make request to API
        const Http = new XMLHttpRequest();
        Http.open("GET", "/search?text=" + payload);
        Http.send();
        Http.onreadystatechange = function(){
            if (this.readyState == 4 && this.status == 200) {
                // when done and successful
                resp = JSON.parse(this.responseText);
                items = ""

                for (var i = 0; i < resp.items.length; i++) {
                    items += "<div class='item'>" + resp.items[i] + "</div>";
                }
                document.getElementById("results").innerHTML = items;

                return;
            }
        };
    };
});