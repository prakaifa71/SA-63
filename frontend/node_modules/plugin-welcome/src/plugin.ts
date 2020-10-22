import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/Users';
import login from './components/Login';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/user', CreateUser);
    router.registerRoute('/login', login);
  },
});
 

