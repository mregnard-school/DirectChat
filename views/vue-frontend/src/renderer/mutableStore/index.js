const store = {
  state: {},
  push: (key) => {
    store.state = {
      ...store.state,
      ...key,
    }
  }
};

export default store;