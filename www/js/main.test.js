require("@babel/polyfill");

describe('Show Local Time', () => {

  // Setup dummy dom
  document.body.innerHTML = '<div id="remote-tz"></div>'

  // Mock Remote Response
  const expectedTimezone = "UTC";
  const mockResponse = Promise.resolve({
    json: () => Promise.resolve({ "timezone": expectedTimezone }),
  });
  global.fetch = jest.fn(() => Promise.resolve(mockResponse));


  // Run main
  require('./main');

  // Assert Timezone
  it('check remote timezone', () => {
    expect(window.document.getElementById("remote-tz").innerHTML).toBe(expectedTimezone);
  });
});
