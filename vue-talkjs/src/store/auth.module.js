
const getUserFromLocalStorage = () => {
  const user = JSON.parse(localStorage.getItem('user'));
  return user ? { status: { loggedIn: true }, user:user } : { status: { loggedIn: false }, user: null };
};

const initialState = getUserFromLocalStorage();

export const auth = {
  namespaced: true,
  state: initialState,
  actions: {
    login({ commit }, userData) {
      localStorage.setItem('user', JSON.stringify(userData));
      commit('loginSuccess', userData);
    },
    logout({ commit }) {
      localStorage.removeItem('user');
      commit('logout');
    }
  },
  mutations: {
    loginSuccess(state, userData) {
      state.status.loggedIn = true;
      state.user = userData;
    },
    loginFailure(state) {
      state.status.loggedIn = false;
      state.user = null;
    },
    logout(state) {
      state.status.loggedIn = false;
      state.user = null;
    },
  }
};

// Listen for changes in local storage and update the store accordingly
window.addEventListener('storage', (event) => {
  if (event.key === 'user') {
    const user = JSON.parse(event.newValue);
    auth.state = getUserFromLocalStorage();
  }
});
