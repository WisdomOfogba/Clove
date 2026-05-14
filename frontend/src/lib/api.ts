export const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

// Mock Data
export const mockRestaurants = [
  {
    id: 'r1',
    name: 'Legacy Bukka',
    cuisine: 'Traditional',
    location: 'Surulere, Lagos',
    rating: 4.8,
    trustScore: 99.8,
    isVerified: true,
    image: 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&q=80',
    badges: ['HYGIENE MASTER', 'FASTEST'],
    status: 'open',
    menu: [
      { id: 'm1', name: 'Amala & Ewedu', price: 2500, description: 'Hot amala with assorted meat.', image: 'https://images.unsplash.com/photo-1512621776951-a57141f2eefd?auto=format&fit=crop&q=60' },
      { id: 'm2', name: 'Jollof Rice', price: 3000, description: 'Smokey party jollof with chicken.', image: 'https://images.unsplash.com/photo-1565299624946-b28f40a0ae38?auto=format&fit=crop&q=60' }
    ]
  },
  {
    id: 'r2',
    name: 'The Jollof Loft',
    cuisine: 'Fusion',
    location: 'Ikeja, Lagos',
    rating: 4.6,
    trustScore: 98.4,
    isVerified: true,
    image: 'https://images.unsplash.com/photo-1565299624946-b28f40a0ae38?auto=format&fit=crop&q=80',
    badges: ['TOP RATED', 'HAND-PICKED'],
    status: 'open',
    menu: [
      { id: 'm3', name: 'Seafood Jollof', price: 4500, description: 'Jollof cooked with shrimps and calamari.', image: 'https://images.unsplash.com/photo-1565299624946-b28f40a0ae38?auto=format&fit=crop&q=60' }
    ]
  },
  {
    id: 'r3',
    name: 'Naija Grills',
    cuisine: 'BBQ',
    location: 'Lekki Phase 1, Lagos',
    rating: 4.5,
    trustScore: 97.9,
    isVerified: true,
    image: 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&q=80',
    badges: ['AUDITED TODAY'],
    status: 'open',
    menu: [
      { id: 'm4', name: 'Chicken Suya', price: 3500, description: 'Spicy grilled chicken with onions.', image: 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&q=60' }
    ]
  }
];

export const mockOrders = [
  { id: 'ord_1', restaurantId: 'r1', restaurantName: 'Legacy Bukka', items: [{ name: 'Amala & Ewedu', qty: 2 }], total: 5000, status: 'delivering', date: new Date().toISOString() },
  { id: 'ord_2', restaurantId: 'r2', restaurantName: 'The Jollof Loft', items: [{ name: 'Seafood Jollof', qty: 1 }], total: 4500, status: 'delivered', date: new Date(Date.now() - 86400000).toISOString() },
];

export const mockVendors = [
  { id: 'v1', name: 'Iya Basira', businessName: 'Basira Foods', status: 'verified', trustScore: 99.1, dateJoined: '2024-01-15' },
  { id: 'v2', name: 'Chef Tolu', businessName: 'Tolu Grills', status: 'pending', trustScore: null, dateJoined: '2024-05-12' },
  { id: 'v3', name: 'Obi', businessName: 'Obi Exquisite', status: 'flagged', trustScore: 65.4, dateJoined: '2024-03-10' },
];

// API Simulation
export const api = {
  // Restaurants
  getRestaurants: async () => {
    await delay(600);
    return mockRestaurants;
  },
  getRestaurant: async (id: string) => {
    await delay(400);
    return mockRestaurants.find(r => r.id === id) || null;
  },

  // Auth
  login: async (credentials: any) => {
    await delay(800);
    return { token: 'mock_token', user: { id: 'u1', role: credentials.role || 'customer', name: 'Test User' } };
  },
  register: async (data: any) => {
    await delay(1000);
    return { success: true };
  },

  // Orders
  getOrders: async () => {
    await delay(500);
    return mockOrders;
  },
  getOrder: async (id: string) => {
    await delay(400);
    return mockOrders.find(o => o.id === id) || { id, status: 'processing', total: 2000, items: [] };
  },
  createOrder: async (data: any) => {
    await delay(1200);
    return { success: true, orderId: `ord_${Math.floor(Math.random() * 1000)}` };
  },

  // Vendor
  submitVendorOnboarding: async (data: any) => {
    await delay(2000);
    // Simulate AI audit results
    const score = Math.floor(Math.random() * 20) + 80; // 80 - 100
    return { success: true, trustScore: score, status: score > 90 ? 'approved' : 'flagged' };
  },
  getVendorProfile: async () => {
    await delay(300);
    return mockVendors[0];
  },
  getVendorOrders: async () => {
    await delay(400);
    return mockOrders;
  },

  // Admin
  getAdminStats: async () => {
    await delay(400);
    return {
      totalVendors: 145,
      activeAudits: 12,
      flagged: 3,
      fraudPrevented: '$12,500'
    };
  },
  getAllVendors: async () => {
    await delay(500);
    return mockVendors;
  }
};
