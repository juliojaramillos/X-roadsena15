package co.gov.ciudadania.repositories;

import co.gov.ciudadania.models.Citizen;
import co.gov.ciudadania.models.DocumentType;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface CitizenRepository extends JpaRepository<Citizen, Integer>{

    @Query( "SELECT c FROM Citizen c " +
            "   WHERE c.document.number like :documentNumber" +
            "     AND c.document.documentType = :documentType" )
    Citizen findByDocument( String documentNumber, DocumentType documentType );

}
