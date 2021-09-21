//Requires
const dataBase = require('../../../../models/');

//Service
const getHistorialBehaviorService = async(type, document) => {

  try{

    if(type && document){

      let response;

      const firstUser = await dataBase.Historical.findOne({
        attributes: ['id', 'hasRecord', 'recordType', 'record'],
        where: {document: document, type: type} })
        .then((historics) => {
          if (historics === null){
            response = {status: 200, data: []};
          }else{
            response = {status: 200, data: historics};
          }
        }).catch((err) => {
          //console.log("Error", err);
          return err
        });
      
      return response;
      
    }; 

  } catch (e){

    throw e;

  };

};

module.exports = {
  getHistorialBehaviorService
};