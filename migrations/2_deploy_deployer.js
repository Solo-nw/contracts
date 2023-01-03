var Deployer = artifacts.require("Deployer");

module.exports = function(deployer) {
    deployer.deploy(Deployer);
    // Additional contracts can be deployed here
};