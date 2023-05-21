import BottomTabNavigator from './lib/features/bottomTabNavigator';
import './lib/core/firebase/firebase';
import { AuthProvider } from './lib/features/login/parts/authContext';

export default function App() {
  return (
    <AuthProvider>
      <BottomTabNavigator />
    </AuthProvider>
  );
}
