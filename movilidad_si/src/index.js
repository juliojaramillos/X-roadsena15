//Requires
const express = require('express');
const morgan = require('morgan');
const dataBase = require('../models/');
const { api_port, routes } = require('./api/libs/initializers');

//Start
const app = express();

//Settings
app.set('port', process.env.PORT || api_port);

//Middlewares
app.use(express.json());
app.use(morgan('dev'));
app.all('*', function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "*");
  res.header("Access-Control-Allow-Headers", "*");
  res.header('Access-Control-Allow-Methods', 'GET');
  next();
});

//Controller
app.use(require(routes.license+'/route'));

//Server
app.listen( app.get('port'), () => {
  dataBase.sequelize.sync();
  console.log('Server in port:', api_port);
});