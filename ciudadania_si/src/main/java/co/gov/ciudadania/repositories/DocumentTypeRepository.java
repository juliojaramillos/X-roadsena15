package co.gov.ciudadania.repositories;

import co.gov.ciudadania.models.DocumentType;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DocumentTypeRepository extends JpaRepository<DocumentType, Integer>{

    DocumentType findByCode( String code );

}
