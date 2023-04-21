import BottomTabNavigator from './lib/features/bottomTabNavigator';
import './lib/core/firebase/firebase';
import { RecoilRoot } from 'recoil';

export default function App() {
  return (
    <RecoilRoot>
      <BottomTabNavigator />
    </RecoilRoot>
  );
}
