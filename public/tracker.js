"use strict";
let http = new XMLHttpRequest();
let params = JSON.stringify({
    "name": "John"
});
http.open('POST', "http://localhost:3000/trackerData", true);
http.setRequestHeader('Content-type', 'application/json');
http.onreadystatechange = () => {
    console.log(http.responseText);
};
http.send(params);