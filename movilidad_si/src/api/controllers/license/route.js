//Requires
const express = require('express');
const { body } = require('express-validator');

//Router
const router = express.Router();

//Controller
const { createNewLicense } = require('./controller');

//Route
router.post('/citizen/license/create', [
  body('document', 'Documento inválido o no existe').exists().isLength({ min: 6 }).not().isEmpty(),
  body('runt_registry', 'Registro en el RUNT inválido o no existe').exists().isInt().not().isEmpty(),
  body('DocumentTypeId', 'Tipo de documento inválido o no existe').exists().not().isEmpty(),
  body('exam', 'El examen no es válido o no existe').exists().not().isEmpty()  
]
,createNewLicense);

module.exports = router;