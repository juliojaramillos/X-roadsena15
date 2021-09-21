//Requires
const { validationResult } = require('express-validator');

//Imports
const { getHistorialBehaviorService } = require('../../services/historical-behavior/service.js');

//Controller
const getHistorialBehavior = async (req, res, next) => {

  //Validate input
  const errors = validationResult(req); 

  if (!errors.isEmpty()) {
    res.status(422).json({ message: errors.errors[0].msg });
    return;
  }

  let { type, document } = req.query;

  try{
    
    let result = await getHistorialBehaviorService(type, document);
    
    if(result.status === 200){
      res.status(result.status).json(result.data);
    }else{
      res.status(result.status).json({message: result.message});
    }
    next();
  } catch(e){
    res.status(500).json({message: "No es posible traer la información en este momento."});
  };


};

module.exports = {
  getHistorialBehavior
};