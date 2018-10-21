const store = {
  state: {},
  push: (key) => {
    store.state = {
      ...store.state,
      ...key,
    }
  }
};

console.log(store);

export default store;