package co.gov.ciudadania.models;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.io.Serializable;

@Entity
@Table( name = "document_type" )
public class DocumentType implements Serializable{

    private static final long serialVersionUID = 1L;

    @Id
    private Integer id;

    private String name;

    private String code;

    public static final DocumentType CC = new DocumentType( 1, "Cédula de Ciudadanía", "CC" );
    public static final DocumentType TI = new DocumentType( 2, "Tarjeta de Identidad", "TI" );
    public static final DocumentType CE = new DocumentType( 3, "Cédula de Extranjería", "CE" );
    public static final DocumentType NUIP = new DocumentType( 4, "Número Único de Identificación Personal", "NUIP" );



    public DocumentType( ){ }

    private DocumentType( Integer id, String name, String code ){
        this.id = id;
        this.name = name;
        this.code = code;
    }



    public Integer getId( ){
        return id;
    }

    public String getName( ){
        return name;
    }

    public String getCode( ){
        return code;
    }

    public static DocumentType parseCode( String code ){
        if( CC.getCode( ).equals( code ) ) return CC;
        else if( TI.getCode( ).equals( code ) ) return TI;
        else if( CE.getCode( ).equals( code ) ) return CE;
        else if( NUIP.getCode( ).equals( code ) ) return NUIP;
        else return null;
    }

}
