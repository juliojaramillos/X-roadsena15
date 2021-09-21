package co.gov.ciudadania.utils;

import co.gov.ciudadania.models.Citizen;
import co.gov.ciudadania.pojos.CitizenPojo;


public class CitizenResponseMapping{

    public static CitizenPojo getCitizenResponse( Citizen citizen ){
        if( citizen == null ) return null;
        CitizenPojo citizenPojo = new CitizenPojo( );
        citizenPojo.setDocumentType( citizen.getDocument( ).getDocumentType( ).getName( ) );
        citizenPojo.setDocumentNumber( citizen.getDocument( ).getNumber( ) );
        citizenPojo.setExpeditionDate( SystemDateFormat.dateFormat( citizen.getDocument( ).getExpeditionDate( ) ) );
        citizenPojo.setName( citizen.getName( ) );
        citizenPojo.setLastname( citizen.getLastname( ) );
        citizenPojo.setAddress( citizen.getAddress( ) );
        citizenPojo.setGender( citizen.getGender( ).toString( ) );
        citizenPojo.setAge( citizen.getAge( ) );
        return citizenPojo;
    }

}
