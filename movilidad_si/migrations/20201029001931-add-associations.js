'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    /**
     * Add altering commands here.
     *
     * Example:
     * await queryInterface.createTable('users', { id: Sequelize.INTEGER });
     */
    return queryInterface.addColumn(
      'DrivenLicense', // name of Target model
      'documentTypeId', // name of the key we're adding
      {
        type: Sequelize.UUID,
        references: {
          model: 'DocumentType', // name of Source model
          key: 'id',
        },
        onUpdate: 'CASCADE',
        onDelete: 'SET NULL',
      }
    );    
  },

  down: async (queryInterface, Sequelize) => {
    /**
     * Add reverting commands here.
     *
     * Example:
     * await queryInterface.dropTable('users');
     */
    return queryInterface.removeColumn(
      'DrivenLicense', // name of the Target model
      'documentTypeId' // key we want to remove
    );
  }
};
