//Requires
const express = require('express');
const { query } = require('express-validator');

//Router
const router = express.Router();

//Controller
const { getHistorialBehavior } = require('./controller');

//Route
router.get('/behavioral-history', [
  query('type', 'Tipo inválido').exists(),
  query('document', 'Documento inválido').exists()
], getHistorialBehavior);

module.exports = router;