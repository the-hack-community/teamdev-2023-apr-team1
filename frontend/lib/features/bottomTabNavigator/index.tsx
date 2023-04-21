import { MaterialIcons } from '@expo/vector-icons';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';

import LoginScreen from '../login/page';
import MapScreen from '../map/page';
import PostsScreen from '../posts/page';


const Tab = createBottomTabNavigator();

const BottomTabNavigator = () => {
  return (
    <NavigationContainer>
      <Tab.Navigator>
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
          name='画像'
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
