package co.gov.ciudadania.controllers;

import co.gov.ciudadania.models.Citizen;
import co.gov.ciudadania.models.DocumentType;
import co.gov.ciudadania.pojos.CitizenPojo;
import co.gov.ciudadania.repositories.CitizenRepository;
import co.gov.ciudadania.utils.CitizenResponseMapping;

import org.springframework.data.repository.query.Param;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;


@RestController
public class CitizenController{

    private final CitizenRepository citizenRepository;

    public CitizenController( CitizenRepository citizenRepository ){
        this.citizenRepository = citizenRepository;
    }

    @GetMapping( "/citizen" )
    public ResponseEntity<CitizenPojo> getCitizenByDocumentTypeAndDocumentNumber(
            @Param( "document" ) String document, @Param( "type" ) String type ){
        Citizen citizen = citizenRepository.findByDocument( document, DocumentType.parseCode( type ) );
        return citizen == null ?
                new ResponseEntity<>( null, HttpStatus.NOT_FOUND ) :
                new ResponseEntity<>( CitizenResponseMapping.getCitizenResponse( citizen ), HttpStatus.OK );
    }


}
