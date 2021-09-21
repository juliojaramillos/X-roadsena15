//Requires
const { validationResult } = require('express-validator');

//Imports
const { createNewCitizenLicense } = require('../../services/license/service.js');

//Controller
const createNewLicense = async (req, res, next) => {

  //Validate input
  const errors = validationResult(req); 

  if (!errors.isEmpty()) {
    res.status(400).json({ message: errors.errors[0].msg });
    return;
  }

  let { document, DocumentTypeId, runt_registry, exam } = req.body;

  try{
    
    let result = await createNewCitizenLicense(document, DocumentTypeId, runt_registry, exam);
    
    if(result.status === 201){
      res.status(result.status).json(result.data);
    }else{
      res.status(result.status).json(result.data);
    }
    next();
  } catch(e){
    console.log("Error", e);
    res.status(500).json({message: "No es posible crear la licencia en este momento."});
  };


};

module.exports = {
  createNewLicense
};