import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { useOnboardingStore } from '../../../store';
import { Icon } from '@iconify/react';

const schema = z.object({
  name: z.string().min(3, 'Business name is required'),
  address: z.string().min(10, 'Full address is required'),
  phone: z.string().min(11, 'Valid Nigerian phone number is required'),
  category: z.string().min(1, 'Please select a category'),
});

export function Step1() {
  const { data, updateBusiness, setStep } = useOnboardingStore();
  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: zodResolver(schema),
    defaultValues: data.business,
  });

  const onSubmit = (values: any) => {
    updateBusiness(values);
    setStep(2);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="p-8 space-y-6">
      <div className="space-y-1">
        <h2 className="text-xl font-bold text-dark">Business Information</h2>
        <p className="text-sm text-neutral">Tell us about your kitchen or restaurant.</p>
      </div>

      <div className="space-y-4">
        <div className="space-y-1.5">
          <label className="text-xs font-bold text-neutral uppercase tracking-wider flex items-center gap-2">
            <Icon icon="lucide:building-2" className="w-3.5 h-3.5" /> Business Name
          </label>
          <input
            {...register('name')}
            placeholder="e.g. Mama T's Jollof Palace"
            className="w-full bg-slate-50 border border-slate-200 rounded-xl px-4 py-3 outline-none focus:ring-2 ring-primary/20 focus:border-primary transition-all"
          />
          {errors.name && <p className="text-[10px] text-red-500 font-bold">{errors.name.message as string}</p>}
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="space-y-1.5">
            <label className="text-xs font-bold text-neutral uppercase tracking-wider flex items-center gap-2">
              <Icon icon="lucide:phone" className="w-3.5 h-3.5" /> Phone Number
            </label>
            <input
              {...register('phone')}
              placeholder="080 0000 0000"
              className="w-full bg-slate-50 border border-slate-200 rounded-xl px-4 py-3 outline-none focus:ring-2 ring-primary/20 focus:border-primary transition-all"
            />
            {errors.phone && <p className="text-[10px] text-red-500 font-bold">{errors.phone.message as string}</p>}
          </div>

          <div className="space-y-1.5">
            <label className="text-xs font-bold text-neutral uppercase tracking-wider flex items-center gap-2">
              <Icon icon="lucide:utensils" className="w-3.5 h-3.5" /> Category
            </label>
            <select
              {...register('category')}
              className="w-full bg-slate-50 border border-slate-200 rounded-xl px-4 py-3 outline-none focus:ring-2 ring-primary/20 focus:border-primary transition-all appearance-none"
            >
              <option value="">Select category</option>
              <option value="bukka">Bukka / Local Kitchen</option>
              <option value="restaurant">Casual Dining</option>
              <option value="fastfood">Fast Food</option>
              <option value="bakery">Bakery & Pastries</option>
              <option value="chef">Private Chef</option>
            </select>
            {errors.category && <p className="text-[10px] text-red-500 font-bold">{errors.category.message as string}</p>}
          </div>
        </div>

        <div className="space-y-1.5">
          <label className="text-xs font-bold text-neutral uppercase tracking-wider flex items-center gap-2">
            <Icon icon="lucide:map-pin" className="w-3.5 h-3.5" /> Physical Address
          </label>
          <textarea
            {...register('address')}
            placeholder="No. 123 Victoria Island, Lagos"
            rows={3}
            className="w-full bg-slate-50 border border-slate-200 rounded-xl px-4 py-3 outline-none focus:ring-2 ring-primary/20 focus:border-primary transition-all resize-none"
          />
          {errors.address && <p className="text-[10px] text-red-500 font-bold">{errors.address.message as string}</p>}
        </div>
      </div>

      <button
        type="submit"
        className="w-full bg-dark text-white rounded-xl py-4 font-bold hover:opacity-90 transition-all flex items-center justify-center gap-2 shadow-lg"
      >
        Continue to Legal Documents
        <Icon icon="lucide:arrow-right" className="w-5 h-5" />
      </button>
    </form>
  );
}
