const { Model } = require('sequelize');

module.exports = (sequelize, DataTypes) => {
  
  class Historical extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
    }
  };
  
  Historical.init({
    document: {
      type: DataTypes.STRING, unique: true
    },
    type: DataTypes.STRING,
    hasRecord: DataTypes.BOOLEAN,
    recordType: DataTypes.STRING,
    record: DataTypes.STRING
  }, {
    sequelize,
    modelName: 'Historical',
  });

  Historical.sync().then(() => {
    
    Historical.create({
      document: '11036516300',
      type: '1',
      hasRecord: true,
      recordType: 'Negativo',
      record: 'Licencia suspendida',
      createdAt: new Date().toDateString(),
      updatedAt: new Date().toDateString()
    });
    
    Historical.create({
      document: '15129143371',
      type: '4',
      hasRecord: false,
      recordType: null,
      record: null,
      createdAt: new Date().toDateString(),
      updatedAt: new Date().toDateString()
    });

    Historical.create({
      document: '16337688',
      type: '2',
      hasRecord: false,
      recordType: null,
      record: null,
      createdAt: new Date().toDateString(),
      updatedAt: new Date().toDateString()
    });

    Historical.create({
      document: '10900945',
      type: '3',
      hasRecord: true,
      recordType: 'Negativo',
      record: 'Licencia cancelada',
      createdAt: new Date().toDateString(),
      updatedAt: new Date().toDateString()
    });

    Historical.create({
      document: '11773819393',
      type: '1',
      hasRecord: false,
      recordType: null,
      record: null,
      createdAt: new Date().toDateString(),
      updatedAt: new Date().toDateString()
    })

  });

  return Historical;
};