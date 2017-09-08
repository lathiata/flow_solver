var settings = require('./settings.json');

// Configure express setup
var path = require('path');
var bodyParser = require('body-parser');
var express = require('express');
var app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.static(path.join(__dirname, 'web')));

app.get('/', function(req, res) {
    res.sendFile(path.join(__dirname + '/web/index.html'));
});

var listen_on_port = settings.listen_on_port;
app.listen(listen_on_port);
