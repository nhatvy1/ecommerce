import { Injectable, NotFoundException } from '@nestjs/common'
import { CategoryRepository } from './category.repository';
import { ErrorMessage } from 'src/utils/response';

@Injectable()
export class CategoryService {
  constructor(private readonly categoryRepository: CategoryRepository) {}

  async findById(id: string) {
    const category = await this.categoryRepository.findByIdAndSelect(id, ['name'])
    if (!category) throw new NotFoundException(ErrorMessage.ErrNotFound)
  }
}
