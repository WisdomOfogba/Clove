import React, { useEffect, useState } from 'react';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { Link } from 'react-router-dom';
import { api } from '../../lib/api';
import { Badge } from '../../components/ui/Badge';
import { Button } from '../../components/ui/Button';

export function OrdersPage() {
  const [orders, setOrders] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchOrders() {
      const data = await api.getOrders();
      setOrders(data);
      setLoading(false);
    }
    fetchOrders();
  }, []);

  if (loading) {
    return <div className="flex justify-center py-40"><div className="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin"></div></div>;
  }

  return (
    <div className="container mx-auto px-4 py-12 max-w-4xl min-h-[70vh]">
      <h1 className="text-4xl font-black text-dark mb-10">My Orders</h1>

      {orders.length === 0 ? (
        <div className="text-center py-20 bg-slate-50 rounded-[3rem] border border-muted">
          <Icon icon="solar:box-minimalistic-bold-duotone" className="w-20 h-20 text-slate-300 mx-auto mb-4" />
          <h2 className="text-2xl font-black text-dark mb-2">No upcomming bites</h2>
          <p className="text-neutral font-medium mb-6">You haven't placed any orders yet.</p>
          <Link to="/restaurants">
            <Button>Explore Kitchens</Button>
          </Link>
        </div>
      ) : (
        <div className="space-y-6">
          {orders.map((order) => (
            <motion.div 
              key={order.id}
              whileHover={{ y: -2 }}
              className="bg-white p-6 sm:p-8 rounded-[2.5rem] border border-muted shadow-sm hover:shadow-md transition-all sm:flex items-center justify-between gap-6"
            >
              <div className="mb-4 sm:mb-0 space-y-2">
                <div className="flex items-center gap-3 mb-2">
                  <span className="font-bold text-xs uppercase tracking-widest text-neutral">{new Date(order.date).toLocaleDateString()}</span>
                  <Badge variant={order.status === 'delivered' ? 'success' : 'warning'}>
                    {order.status === 'delivering' ? 'On the way' : order.status}
                  </Badge>
                </div>
                <h3 className="text-xl font-black text-dark">{order.restaurantName}</h3>
                <p className="text-sm font-bold text-neutral">
                  {order.items.map((i:any) => `${i.qty}x ${i.name}`).join(', ')}
                </p>
                <div className="text-lg font-black text-primary mt-2">₦{order.total}</div>
              </div>
              
              <div className="flex flex-col sm:items-end gap-3 shrink-0">
                <Link to={`/orders/${order.id}`}>
                  <Button variant={order.status === 'delivered' ? 'outline' : 'primary'} className="w-full sm:w-auto h-12">
                    {order.status === 'delivered' ? 'View Details' : 'Track Live'}
                  </Button>
                </Link>
              </div>
            </motion.div>
          ))}
        </div>
      )}
    </div>
  );
}
