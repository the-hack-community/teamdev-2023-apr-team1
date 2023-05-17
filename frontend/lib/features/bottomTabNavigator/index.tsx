import { MaterialIcons } from '@expo/vector-icons';
import {
  useFonts,
  NotoSansJP_400Regular
} from '@expo-google-fonts/noto-sans-jp';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';

import LoginScreen from '../login/page';
import MapScreen from '../map/page';
import PostsScreen from '../posts/page';
const Tab = createBottomTabNavigator();

const BottomTabNavigator = () => {
  const [fontsLoaded] = useFonts({
    NotoSansJP_400Regular
  });

  if (!fontsLoaded) {
    return (
      <NavigationContainer>
        <></>
      </NavigationContainer>
    );
  }

  return (
    <NavigationContainer>
      <Tab.Navigator
        screenOptions={{
          headerTitleStyle: { fontFamily: 'NotoSansJP_400Regular' },
          tabBarLabelStyle: { fontFamily: 'NotoSansJP_400Regular' }
        }}>
        <Tab.Screen
          name='地図'
          component={MapScreen}
          options={{
            tabBarIcon: ({ color, size }) => (
              <MaterialIcons name='map' size={size} color={color} />
            )
          }}
        />
        <Tab.Screen
          name='投稿画像'
          component={PostsScreen}
          options={{
            tabBarIcon: ({ color, size }) => (
              <MaterialIcons name='photo' size={size} color={color} />
            )
          }}
        />
        <Tab.Screen
          name='ログイン'
          component={LoginScreen}
          options={{
            tabBarIcon: ({ color, size }) => (
              <MaterialIcons name='login' size={size} color={color} />
            )
          }}
        />
      </Tab.Navigator>
    </NavigationContainer>
  );
};

export default BottomTabNavigator;
