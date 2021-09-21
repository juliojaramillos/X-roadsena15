'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class DrivenLicense extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
    }
  };
  DrivenLicense.init({
    document: {
      type: DataTypes.STRING, unique: true
    },
    exam: DataTypes.BOOLEAN,
    runt_registry: DataTypes.STRING
  }, {
    sequelize,
    modelName: 'DrivenLicense',
  });

  DrivenLicense.associate = models => {
    DrivenLicense.hasOne(models.DocumentType);
  };

  return DrivenLicense;
};