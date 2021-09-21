package co.gov.ciudadania.utils;

import java.text.SimpleDateFormat;
import java.util.Date;

public class SystemDateFormat{

    private static final String DEFAULT_DATE_FORMAT = "dd/MM/yyyy";
    private static final java.text.DateFormat SIMPLE_DATE_FORMAT = new SimpleDateFormat( DEFAULT_DATE_FORMAT );

    public static String dateFormat( Date date ){
        return SIMPLE_DATE_FORMAT.format( date );
    }

    public static String dateFormat( Date date, String dateFormatPattern ){
        return new SimpleDateFormat( dateFormatPattern ).format( date );
    }

}
