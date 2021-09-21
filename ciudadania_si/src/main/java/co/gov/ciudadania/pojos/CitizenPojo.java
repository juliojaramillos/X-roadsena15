package co.gov.ciudadania.pojos;

public class CitizenPojo{

    private String documentType;
    private String documentNumber;
    private String expeditionDate;
    private String name;
    private String lastname;
    private String address;
    private String gender;
    private Integer age;

    public String getDocumentType( ){
        return documentType;
    }

    public void setDocumentType( String documentType ){
        this.documentType = documentType;
    }

    public String getDocumentNumber( ){
        return documentNumber;
    }

    public void setDocumentNumber( String documentNumber ){
        this.documentNumber = documentNumber;
    }

    public String getExpeditionDate( ){
        return expeditionDate;
    }

    public void setExpeditionDate( String expeditionDate ){
        this.expeditionDate = expeditionDate;
    }

    public String getName( ){
        return name;
    }

    public void setName( String name ){
        this.name = name;
    }

    public String getLastname( ){
        return lastname;
    }

    public void setLastname( String lastname ){
        this.lastname = lastname;
    }

    public String getAddress( ){
        return address;
    }

    public void setAddress( String address ){
        this.address = address;
    }

    public String getGender( ){
        return gender;
    }

    public void setGender( String gender ){
        this.gender = gender;
    }

    public Integer getAge( ){
        return age;
    }

    public void setAge( Integer age ){
        this.age = age;
    }
}
