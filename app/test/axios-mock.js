const sinon = require('sinon');
const expect = require('chai').expect;
const wrapper = require('../services/axios-wrapper');
const ipList = require('../services/mock').ipList;

describe("Axios Mock Wrapper", () => {
  beforeEach(() => {
    this.sandbox = sinon.createSandbox();
  });
  
  afterEach(() => {
    this.sandbox.restore();
  });
  
  it('should replace axios get with Wrapper', (done) => {
    wrapper.http.get('ipList')
        .then((response) => {
          expect(response).to.have.deep.property('data', ipList);
          done();
        })
        .catch((error) => {
          sinon.assert.fail("", "", error);
        });
  })
});