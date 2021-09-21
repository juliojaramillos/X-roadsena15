//Requires
const dataBase = require('../../../../models');

//Service
const createNewCitizenLicense = async(document, DocumentTypeId, runt_registry, exam) => {

  try{

    let response;

    if(exam === true){
      if(document && DocumentTypeId && String(runt_registry) && exam){

        let response;

        const firstUser = await dataBase.DrivenLicense.create({
          document, exam, runt_registry, DocumentTypeId})
          .then((created) => {
            console.log("Created", created);
            if (created === null){
              response = {status: 204, data: response = []};
            }else{
              response = {status: 201, data: response = {licenseID: created.dataValues.id}};
            }
          }).catch((err) => {
            response = {status: 400, data: []};
            return response;
          });
        
        return response;
        
      };
    }else{

      response = {status: 204, data: "El usuario tiene que haber aprobado el examen."}

      return response;
    }   

  } catch (e){
    console.log("Error", e);
    throw e;

  };

};

module.exports = {
  createNewCitizenLicense
};