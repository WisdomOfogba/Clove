import React, { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { api } from '../../lib/api';
import { Button } from '../../components/ui/Button';

export function OrderTrackingPage() {
  const { id } = useParams();
  const [order, setOrder] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  
  // Simulated tracking live stages
  const [stage, setStage] = useState(1); // 1-4

  useEffect(() => {
    async function loadOrder() {
      if (!id) return;
      const data = await api.getOrder(id);
      setOrder(data);
      setLoading(false);
      
      // progression simulation if not delivered
      if (data?.status !== 'delivered') {
         setStage(2);
         setTimeout(() => setStage(3), 5000);
      } else {
         setStage(4);
      }
    }
    loadOrder();
  }, [id]);

  if (loading) {
    return <div className="flex justify-center py-40"><div className="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin"></div></div>;
  }

  return (
    <div className="min-h-screen bg-slate-50 pb-20">
      
      {/* Simulated Map Header */}
      <div className="h-[40vh] bg-dark relative overflow-hidden flex items-center justify-center">
         <div className="absolute inset-0 opacity-20" style={{ backgroundImage: 'radial-gradient(#94A3B8 1px, transparent 1px)', backgroundSize: '30px 30px' }} />
         <div className="absolute inset-0 bg-gradient-to-t from-slate-50 to-transparent" />
         
         <div className="relative z-10 flex flex-col items-center">
            <Icon icon="solar:map-point-wave-bold-duotone" className="w-16 h-16 text-primary animate-bounce mb-4" />
            <div className="bg-white px-6 py-2 rounded-full shadow-lg border border-muted font-black text-sm">
               Driver is 5 mins away
            </div>
         </div>

         <div className="absolute top-4 left-4 z-20">
            <Link to="/orders" className="w-10 h-10 bg-white shadow-xl rounded-full flex items-center justify-center text-dark hover:bg-slate-100 transition-colors">
              <Icon icon="solar:arrow-left-bold" className="w-5 h-5" />
            </Link>
         </div>
      </div>

      <div className="container mx-auto px-4 max-w-3xl -mt-20 relative z-30">
         <div className="bg-white rounded-[3rem] p-8 sm:p-12 shadow-2xl border border-muted mb-8 text-center sm:text-left sm:flex justify-between items-center gap-8">
            <div>
               <h1 className="text-3xl font-black text-dark mb-1">Order {order.id}</h1>
               <p className="font-bold text-neutral">{order.restaurantName}</p>
            </div>
            <div className="mt-6 sm:mt-0 px-6 py-3 bg-primary/20 text-dark rounded-2xl border border-primary/30 inline-flex items-center gap-3">
               <Icon icon="solar:shield-warning-bold-duotone" className="w-6 h-6 text-primary"/>
               <div className="text-left">
                  <p className="text-[10px] font-black uppercase tracking-widest leading-none">Security Lock</p>
                  <p className="text-sm font-bold mt-1">Geo-fenced Active</p>
               </div>
            </div>
         </div>

         {/* Tracking Steps */}
         <div className="bg-white rounded-[3rem] p-8 sm:p-12 shadow-sm border border-muted">
            <h3 className="text-xl font-black text-dark mb-8">Live Tracking Protocol</h3>
            
            <div className="relative pl-8 space-y-10 border-l-2 border-slate-100">
               {[
                 { step: 1, title: 'AI Kitchen Audit Passed', desc: 'Hygiene confirmed. Prep started.', icon: 'solar:scanner-bold-duotone' },
                 { step: 2, title: 'Geospatially Sealed', desc: 'Bag locked with Clove Seal #8492.', icon: 'solar:lock-password-bold-duotone' },
                 { step: 3, title: 'Courier En Route', desc: 'Rider is strictly following the secure route.', icon: 'solar:routing-2-bold-duotone' },
                 { step: 4, title: 'Delivered Safely', desc: 'Enjoy your verified meal.', icon: 'solar:home-smile-bold-duotone' }
               ].map((s, i) => {
                 const isActive = stage === s.step;
                 const isPast = stage > s.step;
                 return (
                   <motion.div 
                     key={s.step} 
                     className="relative"
                     initial={{ opacity: 0, x: -20 }}
                     animate={{ opacity: 1, x: 0 }}
                     transition={{ delay: i * 0.1 }}
                   >
                     {/* Node */}
                     <div className={`absolute -left-[45px] w-8 h-8 rounded-full border-4 flex items-center justify-center bg-white ${isActive ? 'border-primary animate-pulse' : isPast ? 'border-green-500 text-green-500' : 'border-slate-200'}`}>
                        {isPast ? <Icon icon="solar:check-read-bold" className="w-4 h-4" /> : <div className={`w-2 h-2 rounded-full ${isActive ? 'bg-primary' : 'bg-slate-200'}`} />}
                     </div>

                     <div className={`transition-opacity ${isPast || isActive ? 'opacity-100' : 'opacity-40'}`}>
                        <div className="flex items-center gap-3 mb-1">
                           <Icon icon={s.icon} className={`w-6 h-6 ${isActive ? 'text-primary' : 'text-neutral'}`} />
                           <h4 className="text-lg font-black text-dark">{s.title}</h4>
                        </div>
                        <p className="text-sm font-medium text-neutral">{s.desc}</p>
                     </div>
                   </motion.div>
                 )
               })}
            </div>
         </div>
      </div>
    </div>
  );
}
