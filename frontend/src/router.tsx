import { createBrowserRouter, Navigate } from 'react-router-dom';
import App from './App';
import { CustomerHome } from './pages/customer/Home';
import { VendorOnboarding } from './pages/vendor/onboarding/OnboardingFlow';

import { RestaurantsPage } from './pages/customer/Restaurants';
import { RestaurantDetailPage } from './pages/customer/RestaurantDetail';
import { CartPage } from './pages/customer/Cart';
import { CheckoutPage } from './pages/customer/Checkout';
import { OrderConfirmationPage } from './pages/customer/OrderConfirmation';
import { OrdersPage } from './pages/customer/Orders';
import { OrderTrackingPage } from './pages/customer/OrderTracking';

import { LoginPage } from './pages/shared/Login';
import { RegisterPage } from './pages/shared/Register';

import { VendorDashboard } from './pages/vendor/Dashboard';
import { VendorMeals } from './pages/vendor/Meals';
import { VendorOrders } from './pages/vendor/Orders';
import { VendorProfile } from './pages/vendor/Profile';

import { AdminDashboard } from './pages/admin/Dashboard';
import { AdminVendors } from './pages/admin/Vendors';
import { AdminReviewQueue } from './pages/admin/ReviewQueue';
import { AdminAnalytics } from './pages/admin/Analytics';

import {
  AboutPage,
  HowItWorksPage,
  NotFoundPage
} from './pages/Placeholders';
import VDashboard from './pages/vendor/VDashboard';

const LOGGED_IN = true;

export const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      { index: true, element: <CustomerHome /> },
      { path: 'restaurants', element: <RestaurantsPage /> },
      { path: 'restaurants/:id', element: <RestaurantDetailPage /> },
      { path: 'login', element: <LoginPage /> },
      { path: 'register', element: <RegisterPage /> },
      { path: 'cart', element: <CartPage /> },
      { path: 'checkout', element: <CheckoutPage /> },
      { path: 'order-confirmation', element: <OrderConfirmationPage /> },
      { path: 'orders', element: <OrdersPage /> },
      { path: 'orders/:id', element: <OrderTrackingPage /> },
      { path: 'vendor/login', element: <Navigate to="/login" replace /> },
      { path: 'vendor/onboarding', element: <VendorOnboarding /> },
      { path: 'vendor/onboarding/step1', element: <Navigate to="/vendor/onboarding" replace /> },
      { path: 'vendor/onboarding/step2', element: <Navigate to="/vendor/onboarding" replace /> },
      { path: 'vendor/onboarding/step3', element: <Navigate to="/vendor/onboarding" replace /> },
      { path: 'vendor/onboarding/step4', element: <Navigate to="/vendor/onboarding" replace /> },
      { path: 'vendor/onboarding/processing', element: <Navigate to="/vendor/onboarding" replace /> },
      { path: 'vendor/onboarding/result', element: <Navigate to="/vendor/onboarding" replace /> },
      {
        path: 'vendor/',
        element: <VDashboard />,
        children: [
            { index: true, element: LOGGED_IN ? <Navigate to="/vendor/dashboard" replace /> : <Navigate to="/login" replace /> },
            { path: 'dashboard', element: <VendorDashboard /> },
            { path: 'meals', element: <VendorMeals /> },
            { path: 'orders', element: <VendorOrders /> },
            { path: 'profile', element: <VendorProfile /> },
        ]
      },
      { path: 'admin', element: <AdminDashboard /> },
      { path: 'admin/login', element: <Navigate to="/login" replace /> },
      { path: 'admin/dashboard', element: <AdminDashboard /> },
      { path: 'admin/vendors', element: <AdminVendors /> },
      { path: 'admin/vendors/:id', element: <Navigate to="/admin/vendors" replace /> },
      { path: 'admin/review-queue', element: <AdminReviewQueue /> },
      { path: 'admin/analytics', element: <AdminAnalytics /> },
      { path: 'profile', element: <Navigate to="/" replace /> },
      { path: 'about', element: <AboutPage /> },
      { path: 'how-it-works', element: <HowItWorksPage /> },
      { path: '*', element: <NotFoundPage /> },
    ],
  },
]);
