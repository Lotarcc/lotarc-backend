import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  // Method to return sample data
  getExampleData(name: string): string {
    return `Hello${name ? ' ' + name : ''}, this is your example data!`;
  }
}
