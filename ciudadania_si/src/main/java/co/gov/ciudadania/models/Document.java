package co.gov.ciudadania.models;

import javax.persistence.Column;
import javax.persistence.Embeddable;
import javax.persistence.ManyToOne;
import java.io.Serializable;
import java.sql.Date;

@Embeddable
public class Document implements Serializable{

    @Column( name = "document" )
    private String number;
    @ManyToOne
    private DocumentType documentType;
    private Date expeditionDate;


    public Document( ){ }

    public Document( String number, DocumentType documentType, Date expeditionDate ){
        this.number = number;
        this.documentType = documentType;
        this.expeditionDate = expeditionDate;
    }


    public String getNumber( ){
        return number;
    }

    public void setNumber( String number ){
        this.number = number;
    }

    public DocumentType getDocumentType( ){
        return documentType;
    }

    public void setDocumentType( DocumentType documentType ){
        this.documentType = documentType;
    }

    public Date getExpeditionDate( ){
        return expeditionDate;
    }

    public void setExpeditionDate( Date expeditionDate ){
        this.expeditionDate = expeditionDate;
    }
}
